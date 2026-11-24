from datetime import datetime, time, timezone
from zoneinfo import ZoneInfo
from pydantic import BaseModel, Field

class EventPayload(BaseModel):
    # Enforces RFC 3339 string parsing into an aware datetime
    occurred_at: datetime = Field(..., description="Must include offset or Z")

def process_event(payload: EventPayload, target_tz_name: str = "America/Sao_Paulo") -> time:
    # 1. Normalize to UTC instant
    utc_instant = payload.occurred_at.astimezone(timezone.utc)
    
    # 2. Project into target zone to extract wall-clock time safely
    target_tz = ZoneInfo(target_tz_name)
    local_time = utc_instant.astimezone(target_tz).timetz()
    
    return local_time