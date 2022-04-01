import logging
import sqlite3
import pandas as pd
from pathlib import Path
from typing import List


class DatabaseCommiter:
    def __init__(
        self,
        root: str="backend/data",
        csv_name: str="questions.csv",
        db_name: str="questions.db",
    ) -> None:
        self.logger = get_logger()
        self.root = Path(root)
        self.csv_name = Path(csv_name)
        self.db_name = Path(db_name)

        self.df = pd.read_csv(self.root / self.csv_name)

        self.conn = sqlite3.connect(self.root / self.db_name, isolation_level=None)
        self.cursor = self.conn.cursor()

    def insert_data(self, table: str, data: List) -> None:
        self._create_table(table)
        self.cursor.execute(
            f"INSERT INTO {table} (id, year, genre, question, answer, commentary) VALUES (?,?,?,?,?,?)",
            data
        )

    def _create_table(self, table: str) -> None:
        self.cursor.execute(
            f"CREATE TABLE IF NOT EXISTS {table}(id INT PRIMARY KEY, year INT, genre TEXT, question TEXT, answer TEXT, commentary TEXT)",
        )

    def display_db(self, db_name: str, table: str) -> None:
        conn = sqlite3.connect(self.root / db_name)
        df = pd.read_sql(f"SELECT * FROM {table}", conn)
        self.logger.info(df.to_string(index=False))


def get_logger(file: str="backend/out/out.log") -> logging.Logger:
    fmt = logging.Formatter("%(asctime)s :%(name)s: [%(levelname)s]\n%(message)s", "%Y-%m-%dT%H:%M:%S")

    logger = logging.getLogger(__name__)
    logger.setLevel(logging.INFO)

    handler = logging.StreamHandler()
    handler_file = logging.FileHandler(file)

    logger.addHandler(handler)
    logger.addHandler(handler_file)

    handler.setFormatter(fmt)
    handler_file.setFormatter(fmt)
    return logger


if __name__ == "__main__":
    commiter = DatabaseCommiter()

    for _, data in commiter.df.iterrows():
        commiter.insert_data("questions", data.tolist())

    commiter.display_db("questions.db", "questions")
