'use client'
import Header from '../components/Header'
import Form from '../components/Form'
import Sidebar from '../components/Sidebar'
import { Item } from '../entity/item'
import { useState } from 'react'

export default function Home() {
  const [targetItem, setTargetItem] = useState<Item | null>(null)

  const handleTargetItem = (item: Item) => {
    setTargetItem(item)
  }

  return (
    <>
      <Header></Header>
      <main className="flex">
        <Form handleTargetItem={handleTargetItem}></Form>
        <Sidebar targetItem={targetItem}></Sidebar>
      </main>
    </>
  )
}
