GET /api/hrd/employees

{
  "data": [
    {
      "id": 1,
      "nik": "EMP001",
      "name": "Budi",
      "phone": "08123",
      "hire_date": "2024-01-01",
      "status": "active"
    },
    {
      
    }
  ],

}

GET /api/hrd/employees/:id
{
  "id": 1,
  "nik": "EMP001",
  "name": "Budi",
  "address": "Jakarta",
  "status": "active",
  "stores": [
    {
      "store_id": 2,
      "store_name": "Store A",
      "role_at_store": "Kasir"
    }
  ]
}

POST /api/hrd/employees
{
  "nik": "EMP010",
  "name": "Sari",
  "phone": "081234",
  "hire_date": "2025-01-01"
}

PATCH /api/hrd/employees/:id/status
{
  "status": "inactive"
}

POST /api/hrd/store-employees
{
  "store_id": 1,
  "employee_id": 5,
  "role_at_store": "Supervisor"
}

GET /api/hrd/shifts
GET /api/hrd/shifts
POST /api/hrd/shifts
{
  "id": 1,
  "shift_name": "Pagi",
  "start_time": "08:00",
  "end_time": "16:00"
}


POST /api/hrd/work-schedules
{
  "employee_id": 1,
  "store_id": 2,
  "shift_id": 1,
  "date": "2026-02-06"
}

GET /api/hrd/work-schedules/today

{
  "date": "2026-02-06",
  "data": [
    {
      "employee_name": "Budi",
      "store_name": "Store A",
      "shift": "Pagi"
    }
  ]
}

GET /api/hrd/salary-components
[
  { "id": 1, "name": "basic" },
  { "id": 2, "name": "overtime" }
]

POST /api/hrd/employee-salaries
{
  "employee_id": 1,
  "component_id": 1,
  "period": "2026-02-01",
  "amount": 4500000
}

GET /api/hrd/payroll/summary?period=2026-02
{
  "period": "2026-02",
  "total_salary": 125000000,
  "employee_count": 25
}

GET /api/hrd/dashboard/summary
{
  "total_employees": 45,
  "active_employees": 40,
  "inactive_employees": 5,
  "total_stores": 6
}

GET /api/hrd/dashboard/employee-status
[
  { "name": "Active", "value": 40 },
  { "name": "Inactive", "value": 5 }
]

GET /api/hrd/dashboard/employees-by-store
[
  { "store": "Store A", "total": 12 },
  { "store": "Store B", "total": 8 }
]
