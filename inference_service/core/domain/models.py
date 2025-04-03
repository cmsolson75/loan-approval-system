from pydantic import BaseModel, AfterValidator
from typing_extensions import Annotated
from .validators import is_positive, loan_term_validator, credit_score_validator


class InferenceRequest(BaseModel):
    annual_income: Annotated[int, AfterValidator(is_positive)]
    loan_amount: Annotated[int, AfterValidator(is_positive)]
    loan_term: Annotated[int, AfterValidator(loan_term_validator)]
    credit_score: Annotated[int, AfterValidator(credit_score_validator)]


class InferenceResponse(BaseModel):
    approval_status: str
    confidence: float
