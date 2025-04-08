from core.services.approval_service import LoanApprovalService
from repository.onnx_runner import ONNXRunner
from functools import lru_cache
import config


@lru_cache
def load_model_service(
    model_path: str = config.settings.model_path,
    threshold: float = config.settings.prediction_threshold
) -> LoanApprovalService:
    """
    Load and return a singleton LoanApprovalService instance.

    Args:
        model_path (str): Path to the ONNX model.
        threshold (float): Threshold for binary classification.

    Returns:
        LoanApprovalService: Service instance.
    """
    predictor = ONNXRunner(model_path)
    return LoanApprovalService(predictor, threshold)
