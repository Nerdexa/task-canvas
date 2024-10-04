import type { Meta, StoryObj } from '@storybook/react';

import Alert from '.';

const meta = {
  component: Alert,
  tags: ['autodocs'],
} satisfies Meta<typeof Alert>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {};