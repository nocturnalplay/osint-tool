import Head from 'next/head'
import { useState } from 'react'
import Loading from '../../component/loading'
import Link from 'next/link'

export default function Home() {
    const [ip, setip] = useState("")
    const [Ipdata, setIpdata] = useState({ traceroute: "", shodan: "" })
    const [loading, setloading] = useState(false)
    const ipfinder = async () => {
        setloading(true)
        const ipfind = await fetch(`http://192.168.1.7:3333/ip-info?ip=${ip}`).then(res => res.json())
        if (ipfind.status) {
            let iptrack = ipfind.data.traceroute.split("\n")
            let sho = ipfind.data.shodan.split("\n")
            setIpdata(a => ({ ...a, traceroute: iptrack, shodan: sho }))
            setloading(false)
        } else {
            console.log("somthing went wrong")
            setloading(false)
        }
    }
    console.log(Ipdata)

    return (
        <div className='body'>
            <Head>
                <title>CyberCops</title>
                <meta name="description" content="cybercops tools going to fetch the OSINt infermation in the publicly available data" />
            </Head>
            <div className='body-ip'>
                <Link href="/">
                    <h1>Cybercop's</h1>
                </Link>

                <input type="text" placeholder='IP Info' onChange={(e) => { setip(e.target.value) }} value={ip} />
                <input type="button" onClick={ipfinder} value="search" />
            </div>
            {
                loading ?
                    <div className='loading'>
                        <h1>cybercops hunting</h1>
                        <Loading />
                    </div> :
                    <div className='ip-results'>
                        <h1>IP Results</h1>
                        {
                            /* <div className='traceroute-result'>
                                {
                                    Ipdata.traceroute && Ipdata.traceroute.map((a, i) => <div key={a + i}>{a}</div>)
                                }
                            </div> */
                        }
                        <div className='shodan-result'>
                            {
                                Ipdata.shodan && Ipdata.shodan.map((a, i) => <div className='shodan-list' key={a + i}>{a}</div>)
                            }
                        </div>
                    </div>
            }

        </div>
    )
}
