from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    model_path: str = "models/loan_approval.onnx"
    prediction_threshold: float = 0.5

    model_config = SettingsConfigDict(env_file=".env")


settings = Settings()
