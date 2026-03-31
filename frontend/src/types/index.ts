// Auth types
export type AuthUser = {
	id: number
  name: string
	email: string
  sessionToken: string
}

export type AuthCredentials = {
  email: string
  password: string
}

export type RegisterCredentials = {
  name: string
	email: string
	password: string
}

export type AuthResponse = {
	user: AuthUser
	token: string
}

// Job types
export type Job = {
	id: number
	title: string
	description: string
	company?: string
	location?: string
	salary?: string
	ownerId?: number
}

export type Application = {
	id: number
	jobId: number
	userId: number
	status: string
	createdAt: string
	jobTitle?: string
}
