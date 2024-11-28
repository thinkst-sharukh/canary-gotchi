import type { ISequence } from '@/types'
import axios from 'axios'

export const regenerateSequence = ({ id }: { id: string }) => {
  return axios.post<{ data: ISequence }>('/api/regenerate-sequence/' + id)
}

export const validateSequence = ({ id, sequence }: { id: string; sequence: string }) => {
  return axios.post('/api/validate-sequence/' + id, {
    sequence,
  })
}
