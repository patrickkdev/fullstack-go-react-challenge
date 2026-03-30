import axios, { type AxiosError, type AxiosResponse, type InternalAxiosRequestConfig } from 'axios'

const baseURL = import.meta.env.VITE_API_URL ?? 'http://localhost:4000'

export const api = axios.create({
	baseURL,
	headers: {
		'Content-Type': 'application/json',
	},
})

api.interceptors.request.use((config: InternalAxiosRequestConfig) => {
	const token = window.localStorage.getItem('authToken')
	if (token && config.headers) {
		config.headers.Authorization = `Bearer ${token}`
	}
	return config
})

api.interceptors.response.use(
	(response: AxiosResponse) => response,
	(error: AxiosError) => {
		if (error.response?.status === 401) {
			window.localStorage.removeItem('authToken')
		}
		return Promise.reject(error)
	},
)
