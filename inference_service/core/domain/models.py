from pydantic import BaseModel, AfterValidator
from typing_extensions import Annotated
from .validators import is_positive, loan_term_validator, credit_score_validator


class InferenceRequest(BaseModel):
    """
    Request model for inference API.

    Attributes:
        annual_income (int): Annual income in USD.
        loan_amount (int): Requested loan amount.
        loan_term (int): Loan term in years.
        credit_score (int): Credit score of the applicant.
    """
    annual_income: Annotated[int, AfterValidator(is_positive)]
    loan_amount: Annotated[int, AfterValidator(is_positive)]
    loan_term: Annotated[int, AfterValidator(loan_term_validator)]
    credit_score: Annotated[int, AfterValidator(credit_score_validator)]


class InferenceResponse(BaseModel):
    """
    Response model for inference API.

    Attributes:
        approval_status (str): 'approved' or 'rejected'.
        confidence (float): Confidence score in range [0,1].
    """
    approval_status: str
    confidence: float
