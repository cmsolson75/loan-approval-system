from repository.onnx_runner import ONNXRunner


class LoanApprovalService:
    """
    Loan approval service using ONNX model inference.

    Attributes:
        predictor (ONNXRunner): ONNX runtime backend.
        threshold (float): Confidence threshold for approval.
    """
    def __init__(self, predictor: ONNXRunner, threshold: float = 0.5):
        self.predictor = predictor
        self.threshold = threshold

    def evaluate(
        self, annual_income: int, loan_amount: int, loan_term: int, credit_score: int
    ) -> dict:
        """
        Evaluate a loan application using ONNX model.

        Args:
            annual_income (int): Annual income in USD.
            loan_amount (int): Loan amount in USD.
            loan_term (int): Loan term in years.
            credit_score (int): Credit score of the applicant.

        Returns:
            dict: {'status': str, 'confidence': float}
        """
        features = [annual_income, loan_amount, loan_term, credit_score]
        prediction = self.predictor.predict(features)

        status = "approved" if prediction >= self.threshold else "rejected"
        confidence = abs(prediction - 0.5) * 2

        return {"status": status, "confidence": round(confidence, 4)}
