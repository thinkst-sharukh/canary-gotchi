import type { IGotchi } from '@/types'
import axios from 'axios'

export const verifyAuthKey = ({
  token,
  hash,
  id,
  name,
}: {
  token: string
  hash: string
  id: string
  name: string
}) => {
  return axios.post<{ data: IGotchi }>('/api/verify-auth-key', {
    token,
    hash,
    id,
    name,
  })
}

export const existsGotchi = (id: string) => {
  return axios.get<{ data: IGotchi }>('/api/existing-gotchi/' + id)
}
