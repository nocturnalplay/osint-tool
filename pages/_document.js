import { Html, Head, Main, NextScript } from "next/document";

export default function Document() {
    return (
        <Html>
            <Head>
                <meta name="theme-color" content="#97cadb" />
                <script
                    src="https://unpkg.com/@lottiefiles/lottie-player@latest/dist/lottie-player.js"
                    async
                />
            </Head>
            <body>
                <Main />
                <NextScript />
            </body>
        </Html>
    );
}