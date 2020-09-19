/** @jsx jsx */
import React, { ReactElement, Fragment, useState } from "react";
import { TextField } from "@rmwc/textfield";
import { FormField } from "@rmwc/formfield";
import { Typography } from "@rmwc/typography";
import { Select } from "@rmwc/select";
import { Slider } from "@rmwc/slider";
import { Button } from "@rmwc/button";
import { Icon } from "@rmwc/icon";
import { css, jsx } from "@emotion/core";

interface Transformation {
  operations: Operation[];
  sourcePath: string;
  targetPath: string;
}
interface Operation {
  format: number;
  scale: number;
  quality: number;
  height: number;
  width: number;
}

const formField = css`
  margin: 10px 15px;
  & > * {
    margin: 2px 10px;
  }
`;

interface Props {}

function Optimize({}: Props): ReactElement {
  const [transformation, setTransformation] = useState<Transformation>({
    sourcePath: "",
    targetPath: "",
    operations: [],
  });
  return (
    <div className="flex flex-col">
      <FormField>
        <label htmlFor="path">Source Path</label>
        <TextField name="sourcePath" onChange={(e) => setTransformation({...transformation, sourcePath:e.currentTarget.value})}/>
      </FormField>
      <FormField>
        <label htmlFor="path">Target Path</label>
        <TextField name="targetPath" onChange={(e) => setTransformation({...transformation, targetPath:e.currentTarget.value})} />
      </FormField>
      <Typography use="headline4"> Transformations</Typography>

      {transformation.operations.map((operation, index) => (
        <Fragment>
          <section className="flex flex-col m-4">
            <Typography use="headline6">Varient #{index + 1}</Typography>
            <FormField css={formField}>
              <label htmlFor="path">Format: </label>
              <Select
                options={["Keep Original", "AVIF", "WEBP", "JPEG", "PNG"]}
                onChange={(e: any) =>
                  setTransformation({
                    ...transformation,
                    operations: transformation.operations.map((opr) => {
                      if (opr === operation) {
                        opr.format = parseInt(e.target.value);
                        return opr;
                      }
                      return operation;
                    }),
                  })
                }
                defaultValue="JPEG"
              />
            </FormField>
            <FormField css={formField}>
              <label htmlFor="path">Quality: </label>
              <Slider
                value={operation.quality}
                min={1}
                onChange={(evt) => console.log(evt.detail.value)}
                onInput={(evt) => console.log(evt.detail.value)}
                discrete
                step={1}
              />
            </FormField>
            <FormField css={formField}>
              <label htmlFor="path">Height: </label>
              <TextField
                type="number"
                name="height"
                defaultValue={0}
                value={operation.height}
                min={0}
              />
            </FormField>
            <FormField css={formField}>
              <label htmlFor="path">Width: </label>
              <TextField
                type="number"
                name="width"
                defaultValue={0}
                value={operation.width}
                min={0}
              />
            </FormField>
            <FormField css={formField}>
              <label htmlFor="path">Scale: </label>
              <TextField
                type="number"
                name="scale"
                value={operation.scale}
                defaultValue={100}
                max={10000}
              />
            </FormField>
          </section>
          <hr />
        </Fragment>
      ))}
      <Button
        label="Add Transformation"
        onClick={() =>
          setTransformation({
            ...transformation,
            operations: [
              ...transformation.operations,
              { format: 0, height: 0, quality: 80, width: 100, scale: 100 },
            ],
          })
        }
        icon={<Icon icon={{ icon: "add", size: "large" }} />}
      />
      <Button raised onClick={() => console.log(transformation)}>Optimize</Button>
    </div>
  );
}

export default Optimize;
