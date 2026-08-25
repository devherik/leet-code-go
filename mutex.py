import threading
import asyncio
from dataclasses import dataclass, field

class InsufficientFoundError(Exception):
    pass

@dataclass
class SafeBankAccount:
    _balance: float
    _lock: threading.Lock = field(default_factory=threading.Lock, init=False, repr=False)

    def deposit(self, amount: float) -> None:
        with self._lock:
            self._balance += amount
    
    def withdraw(self, amount: float) -> None:
        with self._lock:
            if self._balance < amount:
                raise InsufficientFoundError("Saldo insuficiente")
            self._balance -= amount
    
    @property
    def balance(self) -> float:
        with self._lock:
            return self._balance


class AsyncCache:
    """Cache in memory prtected against Thundering Herd problem"""
    def __init__(self) -> None:
        self._data: dict[str, str] = {}
        self._lock = asyncio.Lock()
    
    async def get_or_set(self, key: str, default_value_coro) -> str:
        # fast reading in caching
        if key in self._data:
            return self._data[key]
        
        # I/O demmand async lock to prevent concurrency
        async with self._lock:
            # double check
            if key in self._data:
                return self._data[key]
            
            compute_value = await default_value_coro()
            self._data[key] = compute_value
            return compute_value