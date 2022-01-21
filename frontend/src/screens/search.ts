import type { GifFetchError } from "../types";

const defaultLimit = 50;

export interface Provider {
    name: string;
    apiKey: string;
    offset: number;
    search: (query: string, resetOffset?: boolean) => Promise<any>;
}

export interface SearchResult {
    preview_url: string;
    original_url: string;
}

export const Giphy: Provider = {
    name: 'Giphy',
    apiKey: 'oTXCaDQKRKtGPpOwRTYVvJjs40mHIygr',
    offset: 0,
    search: async function (query: string, resetOffset = false): Promise<Array<SearchResult>> {

        if (resetOffset) {
            this.offset = 0;
        }

        let res: Response;
        try {
            res = await fetch(`https://api.giphy.com/v1/gifs/search?api_key=${this.apiKey}&q=${encodeURI(query)}&limit=${defaultLimit}&offset=${this.offset}`);
        } catch (e) {
            throw "Network request failed (this should only occur if permissions are missing / endpoint does not get hit)"
        }
        if (!res.ok) {
            let json = await res.json()
            let g: GifFetchError = {
                'statusCode': res.status,
                'statusText': res.statusText,
                'redirected': res.redirected,
                'json': json
            }
            throw g;
        }

        const body: any = await res.json();

        this.offset += 20;

        const results: Array<SearchResult> = [];
        for (const item of body.data) {
            results.push({
                original_url: item.images.original.url,
                preview_url: item.images.downsized.url,
            });
        }

        return results;
    }
};

export const Tenor: Provider = {
    name: 'Tenor',
    apiKey: 'LIDSRZULELA',
    offset: 0,
    search: async function (query: string, resetOffset = false): Promise<Array<SearchResult>> {

        if (resetOffset) {
            this.offset = 0;
        }
        this.lastQuery = query;

        const res: Response = await fetch(`https://g.tenor.com/v1/search?q=${encodeURI(query)}&key=${this.apiKey}&limit=${defaultLimit}&pos=${this.offset}`);
        if (!res.ok) {
            let json = await res.json()
            let g: GifFetchError = {
                'statusCode': res.status,
                'statusText': res.statusText,
                'redirected': res.redirected,
                'json': json
            }
            throw g;
        }
        const body: any = await res.json();

        this.offset = parseInt(body.next);

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
