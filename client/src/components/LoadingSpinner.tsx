import React, { ReactElement } from 'react';
import { Spin } from 'antd';
import { LoadingOutlined } from '@ant-design/icons';

interface Props {
    
}
const antIcon = <LoadingOutlined style={{ fontSize: 24 }} spin />;


export default function LoadingSpinner({}: Props): ReactElement {
    return (
        <Spin indicator={antIcon} />
    )
}
