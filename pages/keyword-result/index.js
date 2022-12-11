import Head from 'next/head'
import { useState } from 'react'
import Loading from '../../component/loading'
import Link from 'next/link'

export default function Home() {
    const [ip, setip] = useState("")
    const [Ipdata, setIpdata] = useState()
    const [loading, setloading] = useState(false)
    const ipfinder = async () => {
        setloading(true)
        const ipfind = await fetch(`http://192.168.1.7:3333/search-key?q=${ip}`).then(res => res.json())
        if (ipfind.status) {
            let keyw = ipfind.data.split("\n")
            let se = keyw.indexOf("https://ahmia.fi/search/?q=weapon&d=30")
            keyw = keyw.splice(se + 1)
            se = keyw.map(a => a.split("redirect_url="))
            console.log("sdfsdf:", se)
            keyw = []
            for (let i = 0; i < se.length; i++) {
                keyw.push(se[i][1])
            }
            setIpdata(keyw)
            setloading(false)
        } else {
            console.log("somthing went wrong")
            setloading(false)
        }
    }

    return (
        <div className='body-dark'>
            <Head>
                <title>CyberCops</title>
                <meta name="description" content="cybercops tools going to fetch the OSINt infermation in the publicly available data" />
            </Head>
            <div className='body-ip'>
                <Link href="/">
                    <h1>Cybercop's</h1>
                </Link>
                <input type="text" placeholder='Darkweb' onChange={(e) => { setip(e.target.value) }} value={ip} />
                <input type="button" onClick={ipfinder} value="search" />
            </div>
            {
                loading ?
                    <div className='loading'>
                        <h1>cybercops hunting</h1>
                        <Loading />
                    </div> :
                    <div className='ip-results'>
                        <h1>Darkweb Results</h1>
                        <div className='shodan-result-dark'>
                            {
                                Ipdata && Ipdata.map((a, i) => a && <div className='keyword-result-list' key={a + i}>{a}</div>)
                            }
                        </div>
                    </div>
            }

        </div>
    )
}
