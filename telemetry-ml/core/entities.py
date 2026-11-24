from dataclasses import dataclass

from pydantic import BaseModel, Field, PositiveFloat, PositiveInt


class TelemetryRecord(BaseModel):
    """System boundary validation contract for incoming raw records."""

    vibration_hz: float = Field(..., ge=0.0, description="Motor vibration frequency")
    rpm: float = Field(..., ge=0.0, description="Rotational speed")
    operating_hours: float = Field(
        ..., ge=0.0, description="Total logged machine hours"
    )
    ambient_temp_c: float = Field(
        ..., ge=-40.0, le=85.0, description="Ambient room temperature"
    )
    target_temperature_c: float = Field(
        ..., ge=-273.15, description="Observed operating temperature"
    )


class TrainingConfig(BaseModel):
    """Immutable hyperparameters defining the numerical optimization constraints."""

    learning_rate: PositiveFloat = Field(
        default=1e-3, description="Optimizer step size"
    )
    epochs: PositiveInt = Field(default=100, description="Full optimization passes")
    batch_size: PositiveInt = Field(
        default=32, description="Mini-batch stochastic sample size"
    )
    l2_penalty: float = Field(
        default=0.0, ge=0.0, description="Weight decay regularization"
    )
    validation_split: float = Field(
        default=0.2, gt=0.0, lt=1.0, description="Holdout ratio"
    )
    random_seed: int = Field(default=42, description="Deterministic run seed")


@dataclass(frozen=True)
class EvaluationMetrics:
    """Read-only value object capturing convergence and generalization performance."""

    loss: float
    mean_absolute_error: float
    r_squared: float
    grad_norm_sanity: bool


@dataclass(frozen=True)
class ModelArtifact:
    """
    Decoupled representation of the learned parameters.
    Does not depend on PyTorch state_dict internals for inspection.
    """

    weights: list[float]
    bias: float
    feature_means: list[float]
    feature_stds: list[float]
    feature_names: list[str]
