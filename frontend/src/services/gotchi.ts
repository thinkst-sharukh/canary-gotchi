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

export const getGotchi = (id: string) => {
  return axios.get<{ data: IGotchi }>('/api/gotchi/' + id)
}

export const updateGotchi = ({ id, hash, token }: { id: string; hash: string; token: string }) => {
  return axios.patch<{ data: IGotchi }>('/api/gotchi/' + id, {
    hash,
    token,
  })
}
