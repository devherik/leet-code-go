export interface CreditCheckPort {
    checkStore(cpf: string): Promise<{ approved: boolean; score: number }>;
}