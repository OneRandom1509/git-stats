

func processRepositories(email string) map[int]int {
    filePath := getDotFilePath()
    repos := parseFileLinesToSlice(filePath)
    daysInMap := daysInLastSixMonths

    commits := make(map[int]int, daysIntMap)
    for i:= daysInMap; i>0; i-- {
        commit[i] = 0
    }

    for _, path := range repos {
        commits = filCommits(email, path, commits)
    }
    return commits
}
