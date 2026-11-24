import StubCreditCheckAdapter from "stub.ts";

describe('Loan Application', () => {
    it('should approve loan when credit score is high', async () => {
        const stubAdapter = new StubCreditCheckAdapter();
        
        const loanService = new LoanService(stubAdapter);
        const result = await loanService.applyForLoan('123.456.789-00', 5000);
        
        expect(result.status).toBe('APPROVED');
    });
});
