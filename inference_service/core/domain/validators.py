def credit_score_validator(score: int) -> int:
    """
    Validate that the credit score is in the range [300, 900].

    Args:
        score (int): Credit score to validate.

    Returns:
        int: Validated credit score.

    Raises:
        ValueError: If score is out of bounds.
    """
    if score > 900 or score < 300:
        raise ValueError(f"Credit score must be between 300 and 900. Got: {score}")
    return score


def loan_term_validator(term: int) -> int:
    """
    Validate that the loan term is between 2 and 30 years.

    Args:
        term (int): Loan term in years.

    Returns:
        int: Validated term.

    Raises:
        ValueError: If term is out of range.
    """
    if term < 2 or term > 30:
        raise ValueError("Loan term must be between 2 and 30 years")
    return term


def is_positive(value: int) -> int:
    """
    Validate that the input value is non-negative.

    Args:
        value (int): Value to validate.

    Returns:
        int: Validated value.

    Raises:
        ValueError: If value is negative.
    """
    if value < 0:
        raise ValueError(f"Can't be negative")
    return value
