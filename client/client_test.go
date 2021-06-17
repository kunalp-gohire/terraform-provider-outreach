package client

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_GetUserData(t *testing.T) {
	testCases := []struct {
		testName     string
		userID       string
		expectErr    bool
		expectedResp *Data
	}{
		{
			testName:  "user exists",
			userID:    "3",
			expectErr: false,
			expectedResp: &Data{
				Data: User{
					Type: "user",
					ID:   3,
					Attributes: Attributes{
						Email:       "kpgkunalgohire@gmail.com",
						FirstName:   "User11",
						LastName:    "Test11",
						Locked:      true,
						UserName:    "Test_User",
						PhoneNumber: "",
						Title:       "Test",
					},
				},
			},
		},
		{
			testName:     "user does not exist",
			userID:       "123",
			expectErr:    true,
			expectedResp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client, _ := NewClient(os.Getenv("OUTREACH_CLIENT_ID"), os.Getenv("OUTREACH_CLIENT_SECRET"), os.Getenv("OUTREACH_REFRESH_TOKEN"),os.Getenv("OUTREACH_REDIRECT_URI"))
			user, err := client.GetUserData(tc.userID)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_CreateUser(t *testing.T) {
	testCases := []struct {
		testName     string
		newItem      Data
		expectedResp *Data
		expectErr    bool
	}{
		{
			testName: "user created successfully",
			newItem: Data{
				Data: User{
					Type: "user",
					Attributes: Attributes{
						Email:       "kpgkunalgohire77@gmail.com",
						FirstName:   "User77",
						LastName:    "Test77",
						Locked:      true,
						Title:       "Test",
					},
				},
			},
			expectedResp: &Data{
				Data: User{
					Type: "user",
					ID:   20,
					Attributes: Attributes{
						Email:       "kpgkunalgohire77@gmail.com",
						FirstName:   "User77",
						LastName:    "Test77",
						Locked:      true,
						UserName:    "User77_Test77",
						PhoneNumber: "",
						Title:       "Test",
					},
				},
			},
			expectErr: false,
		},
		{
			testName: "user already exists",
			newItem: Data{
				Data: User{
					Type: "user",
					Attributes: Attributes{
						Email:       "kpgkunalgohire@gmail.com",
						FirstName:   "User1",
						LastName:    "Test1",
						Locked:      false,
						PhoneNumber: "",
						Title:       "Test",
					},
				},
			},
			expectedResp: &Data{
				Data: User{
					Type: "user",
					ID:   3,
					Attributes: Attributes{
						Email:       "kpgkunalgohire@gmail.com",
						FirstName:   "User1",
						LastName:    "Test1",
						Locked:      false,
						PhoneNumber: "",
						Title:       "Test",
					},
				},
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client, _ := NewClient(os.Getenv("OUTREACH_CLIENT_ID"), os.Getenv("OUTREACH_CLIENT_SECRET"), os.Getenv("OUTREACH_REFRESH_TOKEN"),os.Getenv("OUTREACH_REDIRECT_URI"))
			user, err := client.CreateUser(tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_UpdateUser(t *testing.T) {
	testCases := []struct {
		testName     string
		updatedUser  Data
		expectedResp *Data
		userID       string
		expectErr    bool
	}{
		{
			testName: "user exists",
			updatedUser: Data{
				Data: User{
					Type: "user",
					ID:   3,
					Attributes: Attributes{
						Email:       "kpgkunalgohire@gmail.com",
						FirstName:   "User11",
						LastName:    "Test11",
						Locked:      true,
						UserName:    "Test_User",
						Title:       "Test",
					},
				},
			},
			expectedResp: &Data{
				Data: User{
					Type: "user",
					ID:   3,
					Attributes: Attributes{
						Email:       "kpgkunalgohire@gmail.com",
						FirstName:   "User11",
						LastName:    "Test11",
						Locked:      true,
						UserName:    "Test_User",
						Title:       "Test",
					},
				},
			},
			userID:    "3",
			expectErr: false,
		},
		{
			testName: "item does not exist",
			userID:   "100",
			updatedUser: Data{
				Data: User{
					Type: "user",
					ID:   3,
					Attributes: Attributes{
						Email:       "kpgkunalgohire@gmail.com",
						FirstName:   "User1",
						LastName:    "Test1",
						Locked:      true,
						Title:       "Test",
					},
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client, _ := NewClient(os.Getenv("OUTREACH_CLIENT_ID"), os.Getenv("OUTREACH_CLIENT_SECRET"), os.Getenv("OUTREACH_REFRESH_TOKEN"),os.Getenv("OUTREACH_REDIRECT_URI"))
			_, err := client.UpdateUser(tc.userID, tc.updatedUser)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetUserData(tc.userID)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}
