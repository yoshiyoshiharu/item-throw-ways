'use client'

import { Area } from '../entity/area'
import Calendar from './Calendar'
import React, { useEffect, useState } from 'react'

export default function Sidebar() {
  const [areas, setAreas] = useState<Area[]>([])

  const fetchAreas = async () => {
    const response = await fetch('http://127.0.0.1:3000/areas')
    const data = await response.json()
    setAreas(data)
  }

  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
  }

  useEffect(() => {
    fetchAreas()
  }, [])

  return (
    <div className="border-l-2 w-4/12">
      <div className="flex justify-center mt-6">
        <select onChange={handleChange} className="p-4 border-2 rounded-lg text-gray-600">
          <option>エリアを選択してください</option>
          {
            areas.map((area) => (
              <option key={area.id} value={area.id}>{area.name}</option>
            ))
          }
        </select>
     </div> 
      <Calendar></Calendar>
    </div>
  )
}
