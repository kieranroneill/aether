'use client';
import { Center, Flex, VStack, useDisclosure } from '@chakra-ui/react';
import { FC } from 'react';

// components
import Footer from '@app/components/Footer';
import Header from '@app/components/Header';
import Providers from '@app/components/Providers';

// constants
import { BODY_BACKGROUND_COLOR, DEFAULT_GAP } from '@app/constants';

// fonts
import { latoFont, sawarabiMincho } from '@app/fonts';

// theme
import { theme } from '@app/theme';

// types
import type { ILayoutProps } from '@app/types';
import Navigation from '@app/components/Navigation';

const RootLayout: FC<ILayoutProps> = ({ children }) => {
  const {
    isOpen: isNavigationOpen,
    onClose: onNavigationClose,
    onOpen: onNavigationOpen,
  } = useDisclosure();
  // handlers
  const handleNavigationCloseClick = () => onNavigationClose();
  const handleNavigationOpenClick = () => onNavigationOpen();

  return (
    <html
      className={`${latoFont.variable} ${sawarabiMincho.variable}`}
      lang="en"
    >
      <head>
        <link
          rel="apple-touch-icon"
          sizes="57x57"
          href="/apple-icon-57x57.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="60x60"
          href="/apple-icon-60x60.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="72x72"
          href="/apple-icon-72x72.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="76x76"
          href="/apple-icon-76x76.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="114x114"
          href="/apple-icon-114x114.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="120x120"
          href="/apple-icon-120x120.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="144x144"
          href="/apple-icon-144x144.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="152x152"
          href="/apple-icon-152x152.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="180x180"
          href="/apple-icon-180x180.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="192x192"
          href="/android-icon-192x192.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="32x32"
          href="/favicon-32x32.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="96x96"
          href="/favicon-96x96.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="16x16"
          href="/favicon-16x16.png"
        />
        <link rel="manifest" href="/manifest.json" />

        <meta name="msapplication-TileColor" content="#ffffff" />
        <meta name="msapplication-TileImage" content="/ms-icon-144x144.png" />
        <meta name="theme-color" content="#ffffff" />

        <title>{`${process.env.NEXT_PUBLIC_TAGLINE} | ${process.env.NEXT_PUBLIC_TITLE}`}</title>
      </head>

      <body>
        <Providers theme={theme}>
          {/*navigation menu*/}
          <Navigation
            isOpen={isNavigationOpen}
            onClose={handleNavigationCloseClick}
          />

          <Center as="main" backgroundColor={BODY_BACKGROUND_COLOR}>
            <Flex
              alignItems="center"
              justifyContent="center"
              minH="100vh"
              w="full"
            >
              <VStack alignItems="center" minH="100vh" spacing={0} w="full">
                {/*header*/}
                <Header onNavigationClick={handleNavigationOpenClick} />

                {/*content*/}
                <VStack
                  flexGrow={1}
                  maxW="1000px"
                  pb={DEFAULT_GAP * 2}
                  pt={DEFAULT_GAP}
                  px={DEFAULT_GAP}
                  spacing={0}
                  style={{
                    marginInlineStart: '0px',
                  }}
                  w="full"
                >
                  {children}
                </VStack>

                {/*footer*/}
                <Footer />
              </VStack>
            </Flex>
          </Center>
        </Providers>
      </body>
    </html>
  );
};

export default RootLayout;
