import { Button } from 'antd';
import Title from 'antd/lib/typography/Title';
import React, { MouseEvent, ReactElement } from 'react';

const titleHeader = {
    color:"#777"
}

interface Props {
  title: string;
  value: string;
  actions?: string[];
  onAction?: Function;
}


export default function Stat({ title, value, actions, onAction }: Props): ReactElement {
  const handleAction =(evt:MouseEvent<any>) => {
    const  value= evt.currentTarget.dataset.value;
    onAction && onAction(value);
  }
  
  return (
    <div className="inline-block px-8 py-6">
      <Title level={5} style={titleHeader}>
        {title}
      </Title>
      <Title level={2}>{value}</Title>
      {actions &&
        actions.map((action: string) => {
          return <Button type="link" data-value={action} onClick={handleAction}>{action}</Button>;
        })}
    </div>
  );
}
