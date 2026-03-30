import { Outlet } from "react-router-dom"

interface props {
	User: User
}

export default function DashboardLayout() {
	return (
		<div class="dashboard-layout">
			Olá, { user.name }
		</div>
	)
}
