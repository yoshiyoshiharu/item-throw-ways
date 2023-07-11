'use client'

import { Area } from '../entity/area'
import Calendar from './Calendar'
import React, { useEffect, useState } from 'react'
import { AreaCollectionDate } from '../entity/area_collection_date'
import { Item } from '../entity/item'
import { Kind } from '../entity/kind'

export default function Sidebar({ targetItem }: { targetItem: Item | null }) {
  const [areas, setAreas] = useState<Area[]>([])
  const [areaCollectionDates, setAreaCollectionDates] = useState<AreaCollectionDate[]>([])

  const kindLabel = (kind: Kind) => {
    switch (kind.name) {
      case '可燃ごみ':
        return <span className="ml-2 text-sm text-white bg-red-500 rounded px-2 py-1">{kind.name}</span>
      case '不燃ごみ':
        return <span className="ml-2 text-sm text-white bg-blue-600 rounded px-2 py-1">{kind.name}</span>
      case '資源':
        return <span className="ml-2 text-sm text-white bg-green-700 rounded px-2 py-1">{kind.name}</span>
      case '粗大ごみ':
        return <span className="ml-2 text-sm text-white bg-amber-800 rounded px-2 py-1">{kind.name}</span>
    }
  }

  const fetchAreas = async () => {
    const response = await fetch('http://127.0.0.1:3000/areas')
    const data = await response.json()
    setAreas(data)
  }

  const fetchAreaCollectionDates = async (area_id: string) => {
    const res = await fetch('http://127.0.0.1:3000/area_collect_dates?area_id=' + area_id)
    const data = await res.json()
    setAreaCollectionDates(data)
  }

  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const area_id = e.target.value
    fetchAreaCollectionDates(area_id)
  }

  useEffect(() => {
    fetchAreas()
  }, [])

  return (
    <>
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

        <div className="m-4">
          <Calendar areaCollectionDates={areaCollectionDates} />
        </div>

        {
          targetItem &&
          <div className="m-4">
            <div className="flex items-center justify-center p-4 border-2 rounded-lg shadow-lg shadow-black-500/40">
              <p className="text-md font-bold">{targetItem.name}</p>
              {
                targetItem.kinds.map((kind: Kind) => (
                  kindLabel(kind)
                ))
              }
            </div>
          </div>
        }
      </div>
    </>
  )
}
