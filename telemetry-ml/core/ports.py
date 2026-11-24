from pathlib import Path
from typing import Any, Protocol

from core.entities import (
    EvaluationMetrics,
    ModelArtifact,
    TrainingConfig,
)


class PreparedData(Protocol):
    """
    Opaque data container for model consumption.
    Allows PyTorch tensors inside adapters without coupling domain contracts to torch.Tensor.
    """

    train_features: Any
    train_targets: Any
    val_features: Any
    val_targets: Any
    feature_names: list[str]


class IDataPipeline(Protocol):
    """Boundary responsible for ingestion, schema validation, and leak-free normalization."""

    def prepare_from_source(
        self, source_path: Path, config: TrainingConfig
    ) -> PreparedData:
        """
        Ingests data, validates rows via TelemetryRecord, performs deterministic train/val
        splitting, computes scaling statistics EXCLUSIVELY from train, and returns PreparedData.
        """
        ...


class IModelEngine(Protocol):
    """Boundary responsible for numerical optimization and autograd operations."""

    def train(
        self, data: PreparedData, config: TrainingConfig
    ) -> tuple[ModelArtifact, list[float]]:
        """
        Executes explicit training loop:
        Forward -> Loss -> Zero Grad -> Autograd Backward -> Optimizer Step.
        Returns the finalized ModelArtifact and epoch loss history.
        """
        ...

    def evaluate(
        self, data: PreparedData, artifact: ModelArtifact
    ) -> EvaluationMetrics:
        """Evaluates generalization without tracking computational graphs (torch.no_grad)."""
        ...


class IArtifactRepository(Protocol):
    """Boundary handling serialization and persistence."""

    def save(self, artifact: ModelArtifact, destination: Path) -> None: ...

    def load(self, source: Path) -> ModelArtifact: ...
