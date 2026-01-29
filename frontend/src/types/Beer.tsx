export type Beer = {
    id: number;
    name: string;
    style: string;
    breweryid: number;
    brewery: string;    
    abv: number;
    minibu: number;
    maxibu: number;
};

export type BeerAddResponse = {
    success: boolean;
    beer_id?: number;
    message?: string;
}