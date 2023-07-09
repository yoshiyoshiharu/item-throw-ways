import Header from '../components/Header'
import Form from '../components/Form'
import Sidebar from '../components/Sidebar'

export default function Home() {
  return (
    <>
      <Header></Header>
      <main className="flex">
        <Form></Form>
        <Sidebar></Sidebar>
      </main>
    </>
  )
}
