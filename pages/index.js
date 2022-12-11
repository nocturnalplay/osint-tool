import Head from 'next/head'
import Link from 'next/link'

export default function Home() {

  return (
    <div className='body-main'>
      <Head>
        <title>CyberCops</title>
        <meta name="description" content="cybercops tools going to fetch the OSINt infermation in the publicly available data" />
      </Head>
      <div className='body-main1'>
        <h1>cybercop's</h1>
        <menu>
          <Link href="/ip-result">IP Results</Link>
          <Link href="/domain-result">Domain Results</Link>
          <Link href="/keyword-result">Darkweb Results</Link>
        </menu>
      </div>
    </div>
  )
}
