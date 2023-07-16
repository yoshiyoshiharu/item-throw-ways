'use client'

import { Area } from '../entity/area'
import Calendar from './Calendar'
import React, { useEffect, useState } from 'react'
import { AreaCollectionDate } from '../entity/area_collection_date'
import { Item } from '../entity/item'
import { Kind } from '../entity/kind'

const BASE_URL = process.env.NEXT_PUBLIC_BASE_URL || 'http://127.0.0.1:3000'
const currentDate = new Date();
const currentYear = currentDate.getFullYear();
const currentMonth = currentDate.getMonth() + 1

export default function Sidebar({ targetItem }: { targetItem: Item | null }) {
  const [areas, setAreas] = useState<Area[]>([])
  const [areaCollectionDates, setAreaCollectionDates] = useState<AreaCollectionDate[]>([])

  const kindLabel = (kind: Kind) => {
    switch (kind.name) {
      case '可燃ごみ':
        return <span key={kind.id} className="ml-2 text-sm text-white bg-red-600 rounded px-2 py-1">{kind.name}</span>
      case '不燃ごみ':
        return <span key={kind.id} className="ml-2 text-sm text-white bg-blue-600 rounded px-2 py-1">{kind.name}</span>
      case '資源':
        return <span key={kind.id} className="ml-2 text-sm text-white bg-green-700 rounded px-2 py-1">{kind.name}</span>
      case '粗大ごみ':
        return <span key={kind.id} className="ml-2 text-sm text-white bg-amber-800 rounded px-2 py-1">{kind.name}</span>
    }
  }

  const fetchAreas = async () => {
    const response = await fetch(BASE_URL + '/areas')
    const data = await response.json()
    setAreas(data)
  }

  const fetchAreaCollectionDates = async (area_id: string, year: string, month: string) => {
    const res = await fetch(BASE_URL + '/area_collect_dates?area_id=' + area_id + '&year=' + year + '&month=' + month)
    const data = await res.json()
    setAreaCollectionDates(data)
  }

  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const area_id = e.target.value
    fetchAreaCollectionDates(area_id, currentYear.toString(), currentMonth.toString())
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
            <div className="p-4 border-2 rounded-lg shadow-lg shadow-black-500/40">
              <div className="flex items-center justify-center">
                <p className="text-md font-bold">{targetItem.name}</p>
                {
                  targetItem.kinds.map((kind: Kind) => (
                    kindLabel(kind)
                  ))
                }
              </div>
              {
                targetItem.price != 0 &&
                <div className="text-sm text-gray-700 flex mt-3">
                  <div className="mr-2">料金: </div>
                  <div>{targetItem.price} 円</div>
                </div>
              }
              {
                targetItem.remarks &&
                <div className="text-sm text-gray-700 flex mt-3">
                  <div className="mr-2 whitespace-nowrap">備考: </div>
                  <div>{targetItem.remarks}</div>
                </div>
              }
            </div>
          </div>
        }
      </div>
    </>
  )
}
