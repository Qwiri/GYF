export interface Provider {
    name: string;
    apiKey: string;
    search: (query: string) => Promise<any>;
}

export interface SearchResult {
    preview_url: string;
    original_url: string;
}

export const Giphy: Provider = {
    name: 'Giphy',
    apiKey: 'epo51yrPMWiwryp1w5xbOEK9gJUpGbIX',
    search: async function(query: string): Promise<Array<SearchResult>> {
        const res: Response = await fetch(`https://api.giphy.com/v1/gifs/search?api_key=${this.apiKey}&q=${query}`);
        const body: any = await res.json();
        
        const results: Array<SearchResult> = [];
        for (const item of body.data) {
            results.push({
                original_url: item.images.original.url,
                preview_url: item.images.preview_gif.url,
            });
        }

        return results;
    }
};

export const Tenor: Provider = {
    name: 'Tenor',
    apiKey: 'LIDSRZULELA',
    search: async function(query: string): Promise<Array<SearchResult>> {
        const res: Response = await fetch(`https://g.tenor.com/v1/search?q=${query}&key=${this.apiKey}&limit=50`);
        const body: any = await res.json();
        
        const results: Array<SearchResult> = [];
        for (const item of body.results) {
            results.push({
                original_url: item.media[0].gif.url,
                preview_url: item.media[0].nanogif.url,
            });
        }

        return results;
    }
};

export const Providers: Array<Provider> = [
    Giphy,
    Tenor,
];