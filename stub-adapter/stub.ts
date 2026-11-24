import CreditCheckPort from "interface.ts";

export class StubCreditCheckAdapter implements CreditCeckPort {
    private simulateApproved = true;
    
    public setSimulationMode(approve: boolean) {
        this.simulateApproved = approve;
    }
    
    async checkScore(cpf: string) {
        if (!this.simulateApproved) {
            return { approved: false, score: 350 };
        }
        return { approved: true, score: 850};
    }
}