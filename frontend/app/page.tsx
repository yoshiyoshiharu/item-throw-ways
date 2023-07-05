import Image from 'next/image'

export default function Home() {
  return (
    <>
      <form>
        <input type="text" />
        <button type="submit" className="bg-teal-400 text-white">検索</button>
      </form>
    </>
  )
}
