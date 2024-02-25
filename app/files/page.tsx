'use client';
import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Code,
  Heading,
  HStack,
  IconButton,
  Skeleton,
  Spacer,
  Stack,
  Text,
  Tooltip,
  VStack,
} from '@chakra-ui/react';
import { NextPage } from 'next';
import React, { useState } from 'react';
import { IoCheckmarkDoneOutline, IoDownloadOutline } from 'react-icons/io5';

// components
import CopyIconButton from '@app/components/CopyIconButton';
import VerifyFileProofModal from '@app/components/VerifyFileProofModal';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useButtonHoverBackgroundColor from '@app/hooks/useButtonHoverBackgroundColor';
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';
import useFiles from '@app/hooks/useFiles';
import useSubTextColor from '@app/hooks/useSubTextColor';

// types
import type { IFileResponse } from '@app/types';

// utils
import downloadJSONFile from '@app/utils/downloadJSONFile';

const FilesPage: NextPage = () => {
  // hooks
  const buttonHoverBackgroundColor: string = useButtonHoverBackgroundColor();
  const defaultTextColor: string = useDefaultTextColor();
  const subTextColor: string = useSubTextColor();
  const { files, loading } = useFiles();
  // state
  const [fileToVerify, setFileToVerify] = useState<IFileResponse | null>(null);
  // handlers
  const handleDownloadProofClick = (file: IFileResponse) => () =>
    downloadJSONFile(file.hash, file.proof);
  const handleVerifyFileProofClick = (file: IFileResponse) => () =>
    setFileToVerify(file);
  const handleVerifyFileProofModalClose = () => setFileToVerify(null);
  // renders
  const renderContent = () => {
    let fileKeys: string[];

    if (loading) {
      return Array.from({ length: 3 }, (_, index) => (
        <Skeleton h="20px" key={`files-page-skeleton-item-${index}`} w="full" />
      ));
    }

    if (files) {
      fileKeys = Object.keys(files);

      if (fileKeys.length > 0) {
        return (
          <Accordion allowMultiple={true} w="full">
            {fileKeys.map((key, fileKeyIndex) => (
              <AccordionItem key={`files-page-${key}-${fileKeyIndex}`}>
                {/*accordian button*/}
                <AccordionButton p={DEFAULT_GAP / 2}>
                  <Code
                    color={defaultTextColor}
                    fontSize="md"
                    maxW={650}
                    noOfLines={1}
                    textAlign="left"
                  >
                    {key}
                  </Code>

                  <Spacer />

                  <AccordionIcon />
                </AccordionButton>

                {/*list of files*/}
                <AccordionPanel p={0}>
                  {files[key].map((file, index) => (
                    <HStack
                      alignItems="center"
                      justifyContent="space-between"
                      key={`files-page-${key}-item-${index}`}
                      p={DEFAULT_GAP / 2}
                      spacing={DEFAULT_GAP / 3}
                      w="full"
                    >
                      {/*name*/}
                      <Text
                        color={subTextColor}
                        fontSize="sm"
                        maxW={500}
                        noOfLines={1}
                        textAlign="left"
                      >
                        {file.name}
                      </Text>

                      <Spacer />

                      {/*copy hash button*/}
                      <CopyIconButton
                        ariaLabel={`Copy Hash`}
                        size="md"
                        value={file.hash}
                      />

                      {/*download proof button*/}
                      <Tooltip label={`Download Proof`}>
                        <IconButton
                          _hover={{ bg: buttonHoverBackgroundColor }}
                          aria-label="Download file proof"
                          color={defaultTextColor}
                          icon={<IoDownloadOutline />}
                          onClick={handleDownloadProofClick(file)}
                          size="md"
                          variant="ghost"
                        />
                      </Tooltip>

                      {/*verify proof*/}
                      <Tooltip label={`Verify Proof`}>
                        <IconButton
                          _hover={{ bg: buttonHoverBackgroundColor }}
                          aria-label="Verify proof"
                          color={defaultTextColor}
                          icon={<IoCheckmarkDoneOutline />}
                          onClick={handleVerifyFileProofClick(file)}
                          size="md"
                          variant="ghost"
                        />
                      </Tooltip>
                    </HStack>
                  ))}
                </AccordionPanel>
              </AccordionItem>
            ))}
          </Accordion>
        );
      }
    }

    // when there are no files returned
    return (
      <Stack alignItems="center" flexGrow={1} justify="center" w="full">
        <Text color={defaultTextColor} textAlign="center">
          {`No files found!`}
        </Text>
      </Stack>
    );
  };

  return (
    <>
      <VerifyFileProofModal
        file={fileToVerify}
        onClose={handleVerifyFileProofModalClose}
      />

      <VStack
        alignItems="center"
        justifyContent="flex-start"
        flexGrow={1}
        spacing={DEFAULT_GAP}
        w="full"
      >
        {/*heading*/}
        <Heading color={defaultTextColor} size="lg" textAlign="center" w="full">
          {`Files`}
        </Heading>

        {/*description*/}
        <Text color={defaultTextColor} size="md" textAlign="center" w="full">
          {`Below is a list of files, grouped by their merkle tree roots. You can download a file's proof and use the root you received to verify the file's integrity.`}
        </Text>

        <VStack
          alignItems="center"
          flexGrow={1}
          justify="flex-start"
          spacing={DEFAULT_GAP - 2}
          w="full"
        >
          {renderContent()}
        </VStack>
      </VStack>
    </>
  );
};

export default FilesPage;
