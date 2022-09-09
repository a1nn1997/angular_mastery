import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CommentPayload } from './comment-payload';

@Injectable({
  providedIn: 'root'
})
export class CommentService {
constructor(
    private httpClient: HttpClient
  ) {}
  getAllCommentsForPost(postId: string): Observable<CommentPayload[]> {
    return this.httpClient.get<CommentPayload[]>('http://localhost:8069/api/comment/bypostid/' + postId, {withCredentials: true});
  }

  postComment(commentPayload: CommentPayload): Observable<any> {
    return this.httpClient.post<any>('http://localhost:8069/api/comment/', commentPayload, {withCredentials: true});
  }

  getAllCommentsByUser() {
    return this.httpClient.get<CommentPayload[]>('http://localhost:8069/api/comment/byuser/', {withCredentials: true});
  }
  
}
