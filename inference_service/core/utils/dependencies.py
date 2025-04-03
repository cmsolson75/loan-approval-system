from core.services.approval_service import LoanApprovalService
from repository.onnx_runner import ONNXRunner
from functools import lru_cache
import config


@lru_cache
def load_model_service(
    model_path: str = config.settings.model_path,
    threshold: float = config.settings.prediction_threshold
) -> LoanApprovalService:
    predictor = ONNXRunner(model_path)
    return LoanApprovalService(predictor, threshold)
