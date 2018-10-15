package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

//func (srv *SearchClient) BadClient (req SearchRequest) (*SearchResponse, error) {
//	searcherParams := url.Values{}
//
//	if req.Limit == 100500 {
//		searcherParams.Add("limit", "aa")
//	} else {
//		searcherParams.Add("limit", strconv.Itoa(req.Limit))
//	}
//
//	if req.Offset == 100500 {
//		searcherParams.Add("offset", "aa")
//	} else {
//		searcherParams.Add("offset", strconv.Itoa(req.Offset))
//	}
//	searcherParams.Add("query", req.Query)
//	searcherParams.Add("order_field", req.OrderField)
//
//	if req.OrderBy == 100500 {
//		searcherParams.Add("order_by", "aa")
//	} else {
//		searcherParams.Add("order_by", strconv.Itoa(req.OrderBy))
//	}
//
//	if strings.EqualFold(srv.AccessToken, "badPath") {
//		FilePath = "aaaa.go"
//		srv.AccessToken = "42"
//	}
//
//	if strings.EqualFold(srv.AccessToken, "closeFile") {
//		FilePath = "badRead.xml"
//		srv.AccessToken = "42"
//	}
//
//	if strings.EqualFold(srv.AccessToken, "badXML") {
//		FilePath = "baddataset.xml"
//		srv.AccessToken = "42"
//	}
//
//	searcherReq, _ := http.NewRequest("GET", srv.URL+"?"+searcherParams.Encode(), nil)
//	searcherReq.Header.Add("AccessToken", srv.AccessToken)
//
//	resp, err := client.Do(searcherReq)
//	if err != nil {
//		return nil, fmt.Errorf("unknown error %s", err)
//	}
//
//	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError {
//		return nil, fmt.Errorf("unknown error %s", err)
//	}
//	return nil, nil
//}
//
//func TestLessLimit(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: -1,
//		Offset: 0,
//		Query: "Boyd",
//		OrderField: "id",
//		OrderBy: -1,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil || err.Error() != "limit must be > 0" {
//		t.Errorf("Expected error: limit must be > 0 Got: %#v", err)
//	}
//}
//
//func TestLessOffset(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 30,
//		Offset: -1,
//		Query: "Boyd",
//		OrderField: "id",
//		OrderBy: -1,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil || err.Error() != "offset must be > 0" {
//		t.Errorf("Expected error: offset must be > 0 Got: %#v", err)
//	}
//}
//
//
//func TestTimeout(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "100500",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 10,
//		Offset: 0,
//		Query: "",
//		OrderField: "id",
//		OrderBy: -1,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil || !strings.Contains(err.Error(),"timeout"){
//		t.Errorf("Expected error: timeout Got: %#v", err)
//	}
//}
//
//func TestBadAccessToken(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "4",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 30,
//		Offset: 0,
//		Query: "",
//		OrderField: "id",
//		OrderBy: -1,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil || err.Error() != "Bad AccessToken" {
//		t.Errorf("Expected error: Bad AccessToken Got: %#v", err)
//	}
//}
//
//func TestBadOrderField(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 30,
//		Offset: 0,
//		Query: "",
//		OrderField: "aaa",
//		OrderBy: -1,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil || !strings.Contains(err.Error(), "OrderFeld") {
//		t.Errorf("Expected error: BadOrderFeld Got: %#v", err)
//	}
//}
//
//func TestOrderByIdAsc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 30,
//		Query: "",
//		OrderField: "id",
//		OrderBy: OrderByAsc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestOrderByIdDesc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 30,
//		Query: "",
//		OrderField: "id",
//		OrderBy: OrderByDesc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestOrderByAgeAsc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 30,
//		Query: "",
//		OrderField: "age",
//		OrderBy: OrderByAsc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestOrderByAgeDesc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 30,
//		Query: "",
//		OrderField: "age",
//		OrderBy: OrderByDesc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestOrderByNameAsc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 30,
//		Query: "",
//		OrderField: "name",
//		OrderBy: OrderByAsc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestOrderByNameDesc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 30,
//		Query: "",
//		OrderField: "name",
//		OrderBy: OrderByDesc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestEmptyOrderAsc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 0,
//		Query: "Hilda",
//		OrderField: "",
//		OrderBy: OrderByAsc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestEmptyOrderDesc(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 0,
//		Query: "Hilda",
//		OrderField: "",
//		OrderBy: OrderByDesc,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestEmptyOrderAsIs(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 0,
//		Query: "Hilda",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//func TestLargeOffset(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 25,
//		Offset: 100,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.FindUsers(r)
//	if err != nil {
//		t.Errorf("Unexpected error: %#v", err)
//	}
//}
//
//
func TestOk(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	c := SearchClient{
		AccessToken: "dataset.xml",
		URL: ts.URL,
	}

	r := SearchRequest{
		Limit: 3,
		Offset: 0,
		Query: "",
		OrderField: "",
		OrderBy: OrderByAsIs,
	}

	_, err := c.FindUsers(r)
	if err != nil {
		t.Errorf("Unexpected error: %#v", err)
	}
}

//func TestFatalError(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(BadServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//
//}
//
//func TestBadRequest(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(BadServerJson))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//
//}
//
//func TestBadJson(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(BadJsonServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//
//}
//
//func TestUnknowError(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(BadServerUnknow))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//
//}
//
//func TestNilUrl(t *testing.T) {
//	c := SearchClient{
//		AccessToken: "42",
//		URL: "",
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.FindUsers(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//
//}
//
//func TestBadClientOffset(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 100500,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.BadClient(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//}
//
//func TestBadClientLimit(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 100500,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: OrderByAsIs,
//	}
//
//	_, err := c.BadClient(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//}
//
//func TestBadClientOrderBy(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "42",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: 100500,
//	}
//
//	_, err := c.BadClient(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//}
//
//func TestBadClientWrongPath(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "badPath",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: orderAsc,
//	}
//
//	_, err := c.BadClient(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//}
//
//func TestBadClientBadXML(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	c := SearchClient{
//		AccessToken: "badXML",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: orderAsc,
//	}
//
//	_, err := c.BadClient(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//}
//
//func TestBadClientCloseFile(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
//
//	req := httptest.NewRequest("", ts.URL, nil)
//	req.Header.Add("FilePath", "dataset.xml")
//
//	c := SearchClient{
//		AccessToken: "closeFile",
//		URL: ts.URL,
//	}
//
//	r := SearchRequest{
//		Limit: 3,
//		Offset: 0,
//		Query: "",
//		OrderField: "",
//		OrderBy: orderAsc,
//	}
//
//	_, err := c.BadClient(r)
//	if err == nil {
//		t.Errorf("Expected: error Got: %#v", err)
//	}
//}