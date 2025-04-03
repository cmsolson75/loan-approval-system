from core.utils.dependencies import load_model_service
from core.domain.models import InferenceRequest, InferenceResponse
from core.services.approval_service import LoanApprovalService
from fastapi import Depends, APIRouter, HTTPException

router = APIRouter()

@router.post("/predict", response_model=InferenceResponse)
async def predict(
    request: InferenceRequest,
    model_service: LoanApprovalService = Depends(load_model_service),
) -> InferenceResponse:
    try:
        result = model_service.evaluate(
            request.annual_income,
            request.loan_amount,
            request.loan_term,
            request.credit_score,
        )
    except RuntimeError as e:
        raise HTTPException(status_code=500, detail=str(e))

    return InferenceResponse(
        approval_status=result["status"], confidence=result["confidence"]
    )


# Health
