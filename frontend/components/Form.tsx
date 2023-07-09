'use client'
import React, { useEffect, useState } from 'react'
import { Item } from '../entity/item'

export default function Form() {
  const [allItems, setAllItems] = useState<Item[]>([])
  const [items, setItems] = useState<Item[]>([])
  const [inputValue, setInputValue] = useState('')

  const fetchItems = async () => {
    const response = await fetch('http://127.0.0.1:3000/items')
    const data = await response.json()
    setAllItems(data)
    setItems(data)
  }

  useEffect(() => {
    fetchItems()
  }, [])

  const search = (value: string) => {
    if (value !== '') {
      const filteredList = allItems.filter((item: Item) =>
        item.name.indexOf(value) !== -1
      )
      setItems(filteredList)
      return
    }

    setItems(allItems)
    return
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
    search(e.target.value)
  }

  return (
    <>
      <div className="flex w-full justify-center mt-6">
        <input type="text" value={inputValue} onChange={handleChange} className="block w-10/12 h-10 p-6 border-2 rounded-lg shadow-lg shadow-black-500/40"/>
      </div>

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
