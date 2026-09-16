from dataclasses import dataclass
from datetime import datetime
from typing import List


@dataclass(frozen=True)
class Entry:
    account_id: str
    amount: float
    timestamp: datetime


@dataclass
class Transaction:
    id: str
    entries: List[Entry]

    def is_balance(self) -> bool:
        # Enforce that the sum of all debits and credits equals zero
        return sum(entry.amount for entry in self.entries) == 0


class Ledger:
    def __init__(self):
        self.transactions: List[Transaction] = []

    def record_transaction(self, tx: Transaction) -> None:
        if not tx.is_balance():
            raise ValueError(f"Transaction {tx.id} is unbalanced! Money cannot vanish.")

        self.transactions.append(tx)

    def get_balance(self, account_id: str) -> float:
        balance: float = 0.0

        for tx in self.transactions:
            for entry in tx.entries:
                if entry.account_id == account_id:
                    balance += entry.amount

        return balance


if __name__ == "__main__":
    # Criando as transações
    tx1 = Transaction(
        id="T001",
        entries=[
            Entry(account_id="111", amount=100.0, timestamp=datetime.now()),
            Entry(account_id="222", amount=-100.0, timestamp=datetime.now()),
        ],
    )

    tx2 = Transaction(
        id="T002",
        entries=[
            Entry(account_id="111", amount=50.0, timestamp=datetime.now()),
            Entry(account_id="222", amount=-50.0, timestamp=datetime.now()),
        ],
    )

    # Criando o Ledger
    ledger = Ledger()

    # Adicionando as transações
    ledger.record_transaction(tx1)
    ledger.record_transaction(tx2)

    # Imprimindo os saldos
    print(f"Saldo da conta 111: {ledger.get_balance('111')}")
    print(f"Saldo da conta 222: {ledger.get_balance('222')}")

    # Verificando se as transações estão balanceadas
    print(f"Transação T001 balanceada: {tx1.is_balance()}")
    print(f"Transação T002 balanceada: {tx2.is_balance()}")
