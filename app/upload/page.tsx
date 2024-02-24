'use client';
import {
  Button,
  CreateToastFnReturn,
  Heading,
  HStack,
  Table,
  TableCaption,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tooltip,
  Tr,
  useToast,
  VisuallyHiddenInput,
  VStack,
} from '@chakra-ui/react';
import axios, { AxiosError, AxiosResponse } from 'axios';
import { Metadata, NextPage } from 'next';
import React, { ChangeEvent, MutableRefObject, useRef, useState } from 'react';
import { IoCloudUploadOutline, IoDocumentsOutline } from 'react-icons/io5';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';
import useLogger from '@app/hooks/useLogger';
import usePrimaryColorScheme from '@app/hooks/usePrimaryColorScheme';

// type
import type { ILogger, IUploadResponse } from '@app/types';

// utils
import truncateString from '@app/utils/truncateString';

const UploadPage: NextPage = () => {
  const inputRef: MutableRefObject<HTMLInputElement | null> =
    useRef<HTMLInputElement | null>(null);
  const toast: CreateToastFnReturn = useToast({
    duration: 1000,
    isClosable: true,
    position: 'top',
  });
  // hooks
  const defaultTextColor: string = useDefaultTextColor();
  const logger: ILogger = useLogger();
  const primaryColorScheme: string = usePrimaryColorScheme();
  // state
  const [fileList, setFileList] = useState<FileList | null>(null);
  const [uploading, setUploading] = useState<boolean>(false);
  // misc

  // handlers
  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) =>
    setFileList(event.target.files);
  const handleSelectFilesClick = () => inputRef.current?.click();
  const handleUploadClick = async () => {
    const _functionName: string = 'handleUploadClick';
    const formData: FormData = new FormData();
    let response: AxiosResponse<IUploadResponse>;

    if (!fileList) {
      logger.debug(
        `${UploadPage.displayName}#${_functionName}: no files selected, ignoring`
      );

      return;
    }

    // create the form data
    Array.from(fileList).forEach((file, index) =>
      formData.append(`file-${index}`, file, file.name)
    );

    try {
      response = await axios.post(
        'http://localhost:3000/files/upload',
        formData
      );

      logger.debug(
        `${UploadPage.displayName}#${_functionName}: successfully upload files`,
        response.data
      );
    } catch (error) {
      logger.error(`${UploadPage.displayName}#${_functionName}:`, error);

      if ((error as AxiosError).isAxiosError) {
        toast({
          description: error.message,
          status: 'error',
          title: `HTTP Status: ${(error as AxiosError).status}`,
        });

        return;
      }

      toast({
        description: error.message,
        status: 'error',
        title: 'Failed to upload files',
      });

      return;
    }
  };

  return (
    <VStack
      alignItems="center"
      justifyContent="flex-start"
      flexGrow={1}
      spacing={DEFAULT_GAP + 2}
      w="full"
    >
      <VisuallyHiddenInput
        onChange={handleFileChange}
        multiple={true}
        ref={inputRef}
        type="file"
      />

      {/*heading*/}
      <Heading color={defaultTextColor} size="lg" textAlign="center" w="full">
        {`Upload`}
      </Heading>

      {/*description*/}
      <Text color={defaultTextColor} size="md" textAlign="center" w="full">
        {`You can upload multiple files and you will receive the Merkle Tree root that can be used to verify the integrity of the uploaded files.`}
      </Text>

      {/*buttons*/}
      <HStack
        alignItems="center"
        justifyContent="center"
        spacing={DEFAULT_GAP}
        w="full"
      >
        {/*select files button*/}
        <Button
          colorScheme={primaryColorScheme}
          isDisabled={uploading}
          minW={250}
          onClick={handleSelectFilesClick}
          rightIcon={<IoDocumentsOutline />}
          variant="solid"
        >
          {`Select Files`}
        </Button>

        {/*upload button*/}
        <Button
          colorScheme={primaryColorScheme}
          isDisabled={!fileList}
          isLoading={uploading}
          minW={250}
          onClick={handleUploadClick}
          rightIcon={<IoCloudUploadOutline />}
          variant="solid"
        >
          {`Upload Files`}
        </Button>
      </HStack>

      {/*files to upload table*/}
      <TableContainer w="full">
        <Table>
          {!fileList && <TableCaption>{`No files selected`}</TableCaption>}

          <Thead>
            <Tr>
              <Th>{`Name`}</Th>
              <Th>{`Type`}</Th>
              <Th isNumeric={true}>{`Size`}</Th>
            </Tr>
          </Thead>

          <Tbody>
            {fileList &&
              Array.from(fileList).map(({ name, size, type }, index) => (
                <Tr key={`upload-files-table-item-${index}`}>
                  <Td>
                    <Tooltip label={name}>
                      <Text color={defaultTextColor}>
                        {name.length > 23
                          ? truncateString(name, { end: 10, start: 10 })
                          : name}
                      </Text>
                    </Tooltip>
                  </Td>
                  <Td>
                    <Text color={defaultTextColor}>{type}</Text>
                  </Td>
                  <Td isNumeric={true}>
                    <Text color={defaultTextColor}>{size}</Text>
                  </Td>
                </Tr>
              ))}
          </Tbody>
        </Table>
      </TableContainer>
    </VStack>
  );
};

UploadPage.displayName = 'UploadPage';

export default UploadPage;
