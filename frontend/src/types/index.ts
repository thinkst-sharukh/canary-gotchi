export interface ISequence {
  id: string
  sequence: string
  expires: string
}

export interface IGotchi {
  id: string
  name: string
  level: number
  hash: string
  auth_token: string
  verified: boolean
  sequence: ISequence
}
