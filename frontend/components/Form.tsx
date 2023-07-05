'use client'

export default function Form() {
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
    </>
  )
}
