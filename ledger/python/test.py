import pytest
from datetime import datetime
from ledger import Entry, Transaction, Ledger


def test_successful_transaction_updates_balance():
    ledger: Ledger = Ledger()

    tx: Transaction = Transaction(
        id="tx_valid",
        entries=[
            Entry(
                account_id="main",
                amount=100.0,
                timestamp=datetime.now(),
            ),
            Entry(
                account_id="secondary",
                amount=-100.0,
                timestamp=datetime.now(),
            ),
        ],
    )
    ledger.record_transaction(tx)

    assert ledger.get_balance("main") == 100.0
    assert ledger.get_balance("secondary") == -100.0


def test_multiple_transactions_aggregate_correctly():
    ledger: Ledger = Ledger()

    tx1: Transaction = Transaction(
        id="tx_01",
        entries=[
            Entry(
                account_id="main",
                amount=-100.0,
                timestamp=datetime.now(),
            ),
            Entry(
                account_id="secondary",
                amount=100.0,
                timestamp=datetime.now(),
            ),
        ],
    )

    tx2: Transaction = Transaction(
        id="tx_02",
        entries=[
            Entry(
                account_id="teritary",
                amount=-50.0,
                timestamp=datetime.now(),
            ),
            Entry(
                account_id="main",
                amount=50.0,
                timestamp=datetime.now(),
            ),
        ],
    )

    ledger.record_transaction(tx1)
    ledger.record_transaction(tx2)

    assert ledger.get_balance("main") == -50.0
    assert ledger.get_balance("secondary") == 100.0
    assert ledger.get_balance("teritary") == -50.0


def test_unbalanced_transaction_raises_error():
    ledger: Ledger = Ledger()

    invalid_tx: Transaction = Transaction(
        id="tx_invalid",
        entries=[Entry(account_id="main", amount=100.0, timestamp=datetime.now())],
    )

    with pytest.raises(ValueError, match="is unbalanced!"):
        ledger.record_transaction(invalid_tx)

    assert ledger.get_balance("main") == 0.0
    assert len(ledger.transactions) == 0
