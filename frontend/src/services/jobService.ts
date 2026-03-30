import type { Application, Job } from '../types'
import { api } from './api'

export const jobService = {
	getJobs: async (search?: string) => {
		const url = search ? `/jobs?search=${encodeURIComponent(search)}` : '/jobs'
		const { data } = await api.get<Job[]>(url)
		return data
	},
	createJob: async (job: Pick<Job, 'title' | 'description' | 'company' | 'location' | 'salary'>) => {
		const { data } = await api.post<Job>('/jobs', job)
		return data
	},
	applyJob: async (jobId: number) => {
		const { data } = await api.post<Application>(`/jobs/${jobId}/apply`)
		return data
	},
	getApplications: async () => {
		const { data } = await api.get<Application[]>('/applications')
		return data
	},
	getMyJobs: async () => {
		const { data } = await api.get<Job[]>('/jobs/mine')
		return data
	},
}
