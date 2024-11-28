export interface ISequence {
    id: string,
    sequence: string
    expires: string
}

export interface IGotchi {
    id: string
    name: string
    level: string
    hash: string
    auth_token: string
    sequence: ISequence
}