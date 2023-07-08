'use client'
import React, { useEffect, useState } from 'react'
import { Item } from '../entity/item'

export default function Form() {
  const [items, setItems] = useState<Item[]>([])

  const fetchItems = async () => {
    const response = await fetch('http://127.0.0.1:3000/items')
    console.log(response)
    const data = await response.json()
    console.log(data)
    setItems(data)
    console.log(items)
  }

  useEffect(() => {
    fetchItems()
  }, [])

  const handleSubmit = (event: any) => {
    event.preventDefault()

    alert('検索しました')
  }

  return (
    <>
      <form onSubmit={handleSubmit}>
        <div className="flex w-full justify-center mt-6">
          <input type="text" className="block w-10/12 h-10 p-6 border-2 rounded-lg shadow-lg shadow-black-500/40"/>
          <button type="submit" className="bg-cyan-400 text-white box-border p-3">検索</button>
        </div>
      </form>

      <div className="flex flex-wrap justify-center mt-6">
        {items.map((item) => (
          <div key={item.id} className="w-1/4 h-1/4 p-4">
            <div className="flex flex-col items-center justify-center w-full h-full p-4 border-2 rounded-lg shadow-lg shadow-black-500/40">
              <p className="mt-4 text-lg font-bold">{item.name}</p>
            </div>
          </div>
        ))}
      </div>
    </>
  )
}
