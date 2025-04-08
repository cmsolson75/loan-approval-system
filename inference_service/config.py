from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    """
    Runtime configuration loaded from .env or environment.

    Attributes:
        model_path (str): Path to ONNX model file.
        prediction_threshold (float): Threshold to classify approval.
    """
    model_path: str = "models/loan_approval.onnx"
    prediction_threshold: float = 0.5

    model_config = SettingsConfigDict(env_file=".env")


settings = Settings()
