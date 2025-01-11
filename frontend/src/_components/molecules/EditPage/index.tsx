import { useState } from 'react';

import Button from '@/_components/atoms/Button';
import Box from '@/_components/mui/Box';
import Dialog from '@/_components/mui/Dialog';
import TextFiled, { TextFieldProps } from '@/_components/mui/TextField';

export type EditPageProps = {
  id: string;
  open: boolean;
  initialValue: string;
  onClose: () => void;
  onSave: (id: string, value: string) => void;
  onDelete: (id: string) => void;
};

const EditPage: React.FC<EditPageProps> = ({
  id,
  open,
  onClose,
  initialValue,
  onSave,
  onDelete,
}) => {
  const [value, setValue] = useState<string>(initialValue);

  const handleChangeText: TextFieldProps['onChange'] = (event) => {
    const inputText = event.target.value;
    if (inputText.length > 40) return;
    setValue(inputText);
  };

  const resetValue = () => {
    setValue(initialValue);
  };

  return (
    <Dialog
      open={open}
      onClose={() => {
        onClose();
        resetValue();
      }}
    >
      <Box>
        <TextFiled
          variant="outlined"
          value={value}
          onChange={handleChangeText}
        />
        <Box>
          <Button onClick={() => onSave(id, value)}>変更</Button>
          <Button onClick={() => onDelete(id)}>削除</Button>
        </Box>
      </Box>
    </Dialog>
  );
};

export default EditPage;
