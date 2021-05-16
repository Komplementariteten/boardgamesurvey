import { BoardGame } from './boardgame';

export class SurveyItem {
  public Ref: BoardGame | undefined;
  public Votes: number = 0;
  public Description: string | undefined;
  public Title: string | undefined;
}
