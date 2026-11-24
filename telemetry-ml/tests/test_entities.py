from dataclasses import FrozenInstanceError

import pytest
from pydantic import ValidationError

from core.entities import ModelArtifact, TelemetryRecord, TrainingConfig


def test_telemetry_record_valid():
    record = TelemetryRecord(
        vibration_hz=12.5,
        rpm=1800.0,
        operating_hours=120.0,
        ambient_temp_c=25.0,
        target_temperature_c=65.5,
    )
    assert record.vibration_hz == 12.5
    assert record.rpm == 1800.0


def test_telemetry_record_invalid_ambient_temp():
    with pytest.raises(ValidationError):
        TelemetryRecord(
            vibration_hz=10.0,
            rpm=1500.0,
            operating_hours=10.0,
            ambient_temp_c=90.0,  # exceeds max 85.0
            target_temperature_c=50.0,
        )


def test_training_config_defaults():
    config = TrainingConfig()
    assert config.learning_rate == 1e-3
    assert config.epochs == 100
    assert config.batch_size == 32
    assert config.validation_split == 0.2


def test_training_config_invalid_split():
    with pytest.raises(ValidationError):
        TrainingConfig(validation_split=1.5)


def test_model_artifact_immutable():
    artifact = ModelArtifact(
        weights=[0.1, 0.2],
        bias=0.05,
        feature_means=[1.0, 2.0],
        feature_stds=[0.5, 0.5],
        feature_names=["f1", "f2"],
    )
    with pytest.raises(FrozenInstanceError):
        attr_name = "bias"
        setattr(artifact, attr_name, 1.0)
