import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'

import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import { PrivateRoute, PublicRoute } from './components/RouteGuards'
import { AuthProvider } from './contexts/AuthContext'
import DashboardLayout from './pages/dashboard/index'
import JobsPage from './pages/dashboard/jobs'
import MyApplicationsPage from './pages/dashboard/my-applications'
import MyJobsPage from './pages/dashboard/my-jobs'
import LandingPage from './pages/landing-page'
import Login from './pages/login'
import Register from './pages/register'

const router = createBrowserRouter([
	{
		path: '/',
		element: <LandingPage />,
	},
	{
		path: '/login',
		element: (
			<PublicRoute>
				<Login />
			</PublicRoute>
		),
	},
	{
		path: '/register',
		element: (
			<PublicRoute>
				<Register />
			</PublicRoute>
		),
	},
	{
		path: '/dashboard',
		element: (
			<PrivateRoute>
				<DashboardLayout />
			</PrivateRoute>
		),
		children: [
			{ index: true, element: <JobsPage /> },
			{ path: 'jobs', element: <JobsPage /> },
			{ path: 'my-jobs', element: <MyJobsPage /> },
			{ path: 'my-applications', element: <MyApplicationsPage /> },
		],
	},
])

createRoot(document.getElementById('root')!).render(
	<StrictMode>
		<AuthProvider>
			<RouterProvider router={router} />
		</AuthProvider>
	</StrictMode>,
)


